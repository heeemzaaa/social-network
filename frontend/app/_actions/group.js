"use server"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

/*
    state = {
        error : "for single message error"
        errors : "for form fields errors"
        message : "for success message"
        data : "for returning data (exp grp component) "
    }
*/

export async function addGroupPostAction(prevState, formData) {
    // return state with data containing the post component created to add it dinamically
}

"use server"

// Creates a new group by validating form data and sending it to the group creation API.
// @param {Object} prevState - The previous state of the form, used to preserve state across calls.
// @param {FormData} formData - The form data containing group inputs.
// @returns {Object} - The updated state with errors, success message, or redirect on success.
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
            errors: state.errors,
            error: "Please fix the highlighted fields",
        };
    }

    const newFormData = new FormData();
    newFormData.append('data', JSON.stringify({ title, description }));
    if (img && img.size > 0) {
        newFormData.append('group_img', img);
    }

    try {
        const sessionCookie = cookies().get("session")?.value;
        const res = await fetch(`http://localhost:3000/api/groups/create`, {
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

// Creates a new group event by validating form data and sending it to the event creation API.
// @param {Object} prevState - The previous state of the form, used to preserve state across calls.
// @param {FormData} formData - The form data containing event inputs.
// @returns {Object} - The updated state with errors, success message, or redirect on success.
export async function createGroupEventAction(prevState, formData) {
    const state = {
        errors: {},
        error: null,
        message: null
    }

    const title = formData.get("title")?.trim();
    const description = formData.get("description")?.trim();
    const date = formData.get("date")?.trim();

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
            errors: state.errors,
            error: "Please fix the highlighted fields",
        };
    }

    try {
        const sessionCookie = cookies().get("session")?.value;
        const res = await fetch(`http://localhost:3000/api/events/create`, {
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
