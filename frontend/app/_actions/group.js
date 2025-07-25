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

// Creates a new group by validating form data and sending it to the group creation API Endpoint.
export async function createGroupAction(prevState, formData) {
    const state = {
        errors: {},
        error: null,
        data: null,
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
        }
    }

    const newFormData = new FormData();
    newFormData.append('data', JSON.stringify({ title, description }));
    if (img && img.size > 0) {
        newFormData.append('group', img);
    }

    try {
        const sessionCookie = await cookies().get("session")?.value;
        const res = await fetch(`http://localhost:8080/api/groups/`, {
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
            }
        }

        return {
            ...state,
            message: `${title} Group has been created successfuly.`,
            data
        }
    } catch (error) {
        return {
            ...prevState,
            error: "An unexpected error occurred while creating a new group",
        };
    }
}

export async function createGroupPostAction(prevState, formData) {
    const state = {
        errors: {},
        error: null,
        message: null
    }

    const content = formData.get("content")?.trim();
    const groupId = formData.get("groupId")?.trim();
    const img = formData.get("image");

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
        console.log("post image: ", img)
        newFormData.append('post', img);
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
        console.log("post created: ", data)
        return {
            ...state,
            message: `Post has been created successfuly.`,
            data
        }
    } catch (error) {
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
        data: null,
        message: null
    }

    let title = formData.get("title")?.trim();
    let description = formData.get("description")?.trim();
    let event_date = formData.get("event_date")?.trim();
    let groupId = formData.get("groupId")?.trim();

    if (!title) {
        state.errors.title = "Title is required";
    }

    if (!description) {
        state.errors.description = "Description is required";
    }

    if (!event_date) {
        state.errors.event_date = "Event date is required";
    } else {
        const eventDate = new Date(event_date);
        if (isNaN(eventDate.getTime())) {
            state.errors.event_date = "Invalid date format";
        } else if (eventDate < new Date()) {
            state.errors.event_date = "Event date must be in the future";
        }
    }

    if (Object.keys(state.errors).length > 0) {
        return {
            ...prevState,
            errors: state.errors
        };
    }

    console.log("date before format: ", event_date)
    event_date = formatDate(event_date)
    console.log("date after format: ", event_date)






    try {
        const cookieStore = await cookies();
        const sessionCookie = cookieStore.get("session")?.value;
        const res = await fetch(`http://localhost:8080/api/groups/${groupId}/events/`, {
            method: "POST",
            body: JSON.stringify({ title, description, event_date }),
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
            ...state,
            data,
            message: "New Event created successfully",
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


const formatDate = (dateString) => {
    const date = new Date(dateString);
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0"); // Months are 0-based
    const day = String(date.getDate()).padStart(2, "0");
    const hours = String(date.getHours()).padStart(2, "0");
    const minutes = String(date.getMinutes()).padStart(2, "0");
    const seconds = String(date.getSeconds()).padStart(2, "0");
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}