"use server"

import { cookies } from "next/headers"
import { redirect } from "next/navigation"

export async function loginUser(prevState, formData) {
    const state = {
        errors: {},
        error: null,
        message: null
    }

    const login = formData.get("login")?.trim()
    const password = formData.get("password")?.trim()

    if (!login) state.errors.login = "Field can't be empty"
    if (!password) state.errors.password = "Field can't be empty"

    if (Object.keys(state.errors).length > 0) {
        return state
    }

    try {
        const res = await fetch(`http://localhost:8080/api/auth/login`, {
            method: "POST",
            body: JSON.stringify({ login, password }) // Send credentials
        });

        const data = await res.json();
        if (!res.ok) {
            state.error = data.error || "Login failed";
            state.errors = data.errors || null;
            return state;
        }
        await setCookie(res.headers.get('set-cookie'))
    } catch (error) {
        console.error(error)
        state.error = "An unexpected error occurred";
        return state;
    }
    redirect("/")
}

// Registers a new user by validating form data, handling file uploads, and sending the data to the register API.
// @param {Object} prevState - The previous state of the form, used to preserve state across calls.
// @param {FormData} formData - The form data containing user inputs.
// @returns {Object} - The updated state with errors, success message, or redirect on success.
export async function registerUser(prevState, formData) {
    const state = {
        errors: {},
        error: null,
        message: null
    }

    const email = formData.get("email")?.trim();
    const password = formData.get("password")?.trim();
    const firstname = formData.get("firstname")?.trim();
    const lastname = formData.get("lastname")?.trim();
    const birthdate = formData.get("birthdate")?.trim();
    const nickname = formData.get("nickname")?.trim() || null;
    const aboutMe = formData.get("about_me")?.trim() || null;
    const avatar = formData.get("avatar");

    if (!email) {
        state.errors.email = "Email is required";
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]{2,}$/.test(email)) {
        state.errors.email = "Invalid email format";
    }

    if (!password) {
        state.errors.password = "Password is required";
    } else if (password.length < 6) {
        state.errors.password = "Password must be at least 6 characters";
    }

    if (!firstname) state.errors.firstname = "First name is required";
    if (!lastname) state.errors.lastname = "Last name is required";
    if (!birthdate) {
        state.errors.birthdate = "Date of birth is required";
    } else {
        const dob = new Date(birthdate);
        if (isNaN(dob.getTime())) {
            state.errors.birthdate = "Invalid date";
        }
    }

    if (avatar && avatar.size > 0) {
        const allowedTypes = ["image/jpeg", "image/png", "image/gif"];
        const maxSize = 3 * 1024 * 1024;
        if (!allowedTypes.includes(avatar.type)) {
            state.errors.avatar = "Avatar must be a JPEG, PNG, or GIF image";
        } else if (avatar.size > maxSize) {
            state.errors.avatar = "Avatar file size must be less than 5MB";
        }
    }

    if (Object.keys(state.errors).length > 0) {
        return {
            ...prevState,
            errors: state.errors,
            error: "Please fix the highlighted fields",
        };
    }

    const newFormData = new FormData()
    newFormData.append('data', JSON.stringify(
        { firstname, lastname, birthdate, email, password, nickname, aboutMe }
    ))
    if (avatar && avatar.size > 0) {
        newFormData.append('profile_img', avatar);
    }
    try {
        const res = await fetch(`http://localhost:3000/api/auth/register`, {
            method: "POST",
            body: newFormData,
            credentials: 'include'
        });
        const data = await res.json();
        if (!res.ok) {
            let state = {
                ...prevState,
                error: data.error || "Registeration failed",
                errors: data.errors || null
            }
            return state;
        }
        await setCookie(res.headers.get('set-cookie'))
    } catch (error) {
        console.log(error)
    }
    redirect("/")
}


export async function logout() {
    try {
        const sessionCookie = cookies().get("session")?.value;
        const res = await fetch(`http://localhost:3000/api/auth/logout`, {
            method: "POST",
            credentials: 'include',
            headers: sessionCookie ? { Cookie: `session=${sessionCookie}` } : {}
        });
        if (res.ok) {
            let cookieStore = await cookies()
            cookieStore.delete("session")
        }
    } catch (error) {
        console.log(error)
    }
    redirect("/login")
}


export async function setCookie(cookieStr) {
    const parts = cookieStr.split(';');
    const result = {};

    parts.forEach(part => {
        const trimmed = part.trim();
        if (trimmed.includes('=')) {
            const [key, value] = trimmed.split('=');
            console.log(key)
            if (key === "session") {
                result.name = key
                result.value = value
            } else {
                result[key] = value;
            }
        } else {
            result[trimmed] = true;
        }
    });

    const cookieStore = await cookies()
    cookieStore.set({
        name: result.name,
        value: result.value,
        path: result.Path,
        expires: new Date(result.Expires),
        httpOnly: true,
        sameSite: "lax"
    })
}