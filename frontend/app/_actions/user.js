"use server"

import { redirect } from "next/navigation"

export async function loginUser(prevState, formData) {
    console.info("heeeeere", formData)
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
            body: JSON.stringify() // Send credentials
        });

        const data = await res.json();
        if (!res.ok) {
            state.error = data.error || "Login failed";
            return state;
        }
        // Assuming the backend returns a success message or token
        state.message = data.message || "Login successful";
        redirect("/"); // Redirect on success
    } catch (error) {
        console.error
        state.error = "An unexpected error occurred";
        return state;
    }
}


export async function registerUser(prevState, formData) {
    // validates form data
    // handle the diffrent parts of the form ( one for the image file, and the other for user fields inputs)
    // send formData to the register api and handle the response
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
    const aboutMe = formData.get("aboutMe")?.trim() || null;
    const avatar = formData.get("avatar") || null;

    console.log("this the user avatr: ", avatar)

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
    newFormData.append('profile_img', avatar)
    try {
        const res = await fetch(`http://localhost:8080/api/auth/register`, {
            method: "POST",
            body: newFormData,
            credentials: 'include'
        });
        const data = await res.json();
        if (!res.ok) {
            console.log("data", data)
            let state = {
                ...prevState,
                error: data.error || "Registeration failed",
                errors: data.errors || null
            }
            console.log("state", state)
            return state;
        }
        state.message = data.message || "register successful";
        redirect('/')
    } catch (error) {
        state.error = "An unexpected error occurred: " + error ;
        return state;
    }
}