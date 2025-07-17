"use server"
import { cookies } from "next/headers"

/*
    state = {
        error : "for single message error"
        errors : "for form fields errors"
        message : "for success message"
        data : "for returning data (exp grp component) "
    }
*/

export async function createGroupPostAction(prevState, formData) {
    const state = {
        errors: {},
        error: null,
        message: null
    }

    const content = formData.get("content")?.trim();
    const groupId = formData.get("groupId")?.trim();
    console.log("group id: ", groupId)
    const img = formData.get("img");

    if (!content) {
        state.errors.content = "Content is required";
    }

    if (img && img.size > 0) {
        const allowedTypes = ["image/jpeg", "image/png", "image/gif"];
        const maxSize = 3 * 1024 * 1024; // 3MB
        if (!allowedTypes.includes(img.type)) {
            state.errors.img = "Image must be a JPEG, PNG, or GIF";
        } else if (img.size > maxSize) {
            state.errors.img = "Image file size must be less than 3MB";
        }
    }

    if (Object.keys(state.errors).length > 0) {
        return {
            ...prevState,
            errors: state.errors
        };
    }

    const newFormData = new FormData();
    newFormData.append('data', JSON.stringify({ content }));

    if (img && img.size > 0) {
        newFormData.append('group_img', img);
    }

    try {
        const cookieStore = await cookies();
        const sessionCookie = cookieStore.get("session")?.value;
        const res = await fetch(`http://localhost:8080/api/groups/${groupId}/posts/`, {
            method: "POST",
            body: newFormData,
            credentials: 'include',
            headers: sessionCookie ? { Cookie: `session=${sessionCookie}` } : {}
        })
        const data = await res.json();
        if (!res.ok) {
            return {
                ...prevState,
                error: data.error || "Group creation failed",
                errors: data.errors || null
            };
        }
        return {
            ...prevState,
            message: "Group created successfully",
        }
    } catch (error) {
        return {
            ...prevState,
            error: "An unexpected error occurred",
        };
    }
}

// Creates a new group by validating form data and sending it to the group creation API Endpoint.
export async function createGroupAction(prevState, formData) {
    const state = {
        errors: {},
        error: null,
        message: null
    }

    const title = formData.get("title")?.trim();
    const description = formData.get("description")?.trim();
    const img = formData.get("img");

    if (!title) {
        state.errors.title = "Title is required";
    } else if (title.length < 3) {
        state.errors.title = "Title must be at least 3 characters";
    }

    if (!description) {
        state.errors.description = "Description is required";
    } else if (description.length < 10) {
        state.errors.description = "Description must be at least 10 characters";
    }

    if (img && img.size > 0) {
        const allowedTypes = ["image/jpeg", "image/png", "image/gif"];
        const maxSize = 3 * 1024 * 1024; // 3MB
        if (!allowedTypes.includes(img.type)) {
            state.errors.img = "Image must be a JPEG, PNG, or GIF";
        } else if (img.size > maxSize) {
            state.errors.img = "Image file size must be less than 3MB";
        }
    }

    if (Object.keys(state.errors).length > 0) {
        return {
            ...prevState,
            errors: state.errors
        };
    }

    const newFormData = new FormData();
    newFormData.append('data', JSON.stringify({ title, description }));
    if (img && img.size > 0) {
        newFormData.append('group_img', img);
    }

    try {
        const sessionCookie = await cookies().get("session")?.value;
        const res = await fetch(`http://localhost:8080/api/groups/`, {
            method: "POST",
            body: newFormData,
            credentials: 'include',
            headers: sessionCookie ? { Cookie: `session=${sessionCookie}` } : {}
        });
        const data = await res.json();
        if (!res.ok) {
            return {
                ...prevState,
                error: data.error || "Group creation failed",
                errors: data.errors || null
            };
        }

        return {
            ...prevState,
            message: "Group created successfully",
        };
    } catch (error) {
        console.error(error);
        return {
            ...prevState,
            error: "An unexpected error occurred",
        };
    }
}

// Creates a new group event by validating form data and sending it to the event creation API endpoint.
export async function createGroupEventAction(prevState, formData) {
    const state = {
        errors: {},
        error: null,
        message: null
    }

    const title = formData.get("title")?.trim();
    const description = formData.get("description")?.trim();
    const date = formData.get("date")?.trim();
    const groupId = formData.get("groupId")?.trim();

    if (!title) {
        state.errors.title = "Title is required";
    }

    if (!description) {
        state.errors.description = "Description is required";
    }

    if (!date) {
        state.errors.date = "Event date is required";
    } else {
        const eventDate = new Date(date);
        if (isNaN(eventDate.getTime())) {
            state.errors.date = "Invalid date format";
        } else if (eventDate < new Date()) {
            state.errors.date = "Event date must be in the future";
        }
    }

    if (Object.keys(state.errors).length > 0) {
        return {
            ...prevState,
            errors: state.errors
        };
    }

    try {
        const cookieStore = await cookies();
        const sessionCookie = cookieStore.get("session")?.value;
        const res = await fetch(`http://localhost:8080/api/groups/${groupId}/events/`, {
            method: "POST",
            body: JSON.stringify({ title, description, date }),
            credentials: 'include',
            headers: {
                "Content-Type": "application/json",
                ...(sessionCookie ? { Cookie: `session=${sessionCookie}` } : {})
            }
        });
        const data = await res.json();
        if (!res.ok) {
            console.error(data)
            return {
                ...prevState,
                error: data.error || "Event creation failed",
                errors: data.errors || null
            };
        }
        return {
            ...prevState,
            message: "Event created successfully",
        };
    } catch (error) {
        console.error(error);
        return {
            ...prevState,
            error: "An unexpected error occurred",
        };
    }
}

export async function joinGroupAction(prevState, formData) {

    

}

export async function inviteUsersAction(prevState, formData) {
    const userIds = formData.getAll('userIds');
    const groupId = formData.get('groupId');
    if (userIds.length === 0) {
        return { success: false, message: 'Please select at least one follower to invite.' };
    }
    console.log('Inviting users:', userIds, 'to group:', groupId);
    return { success: true, message: `Invited ${userIds.length} user(s) to group ${groupId}` };
}