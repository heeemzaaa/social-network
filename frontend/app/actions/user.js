"use server"

import { redirect } from "next/navigation"

export async function loginUser(prevState, formData) {
    const state = {
        errors: {},
        error: null,
        message: null
    }

    const username = formData.get("username")?.trim()
    const password = formData.get("password")?.trim()

    if (!username) state.errors.username = "Field can't be empty"
    if (!password) state.errors.password = "Field can't be empty"


    if (Object.keys(state.errors).length > 0) {
        return state
    }

    // here i can send data to backend 

    redirect("/")
}

export async function registerUser(prevState, formData) {
    const state = {
        errors: {},
        error: null,
        message: null
    }


    const email = formData.get("email")?.toString().trim();
    const password = formData.get("password")?.toString().trim();
    const firstname = formData.get("firstname")?.toString().trim();
    const lastname = formData.get("lastname")?.toString().trim();
    const dateOfBirth = formData.get("dateOfBirth")?.toString().trim();
    const nickname = formData.get("nickname")?.toString().trim() || null;
    const aboutMe = formData.get("aboutMe")?.toString().trim() || null;
    const avatar = formData.get("avatar"); // can be a File or null

    if (!email) {
        state.errors.email = "Email is required";
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
        state.errors.email = "Invalid email format";
    }

    if (!password) {
        state.errors.password = "Password is required";
    } else if (password.length < 6) {
        state.errors.password = "Password must be at least 6 characters";
    }

    if (!firstname) state.errors.firstname = "First name is required";
    if (!lastname) state.errors.lastname = "Last name is required";
    if (!dateOfBirth) {
        state.errors.dateOfBirth = "Date of birth is required";
    } else {
        const dob = new Date(dateOfBirth);
        if (isNaN(dob.getTime())) {
            state.errors.dateOfBirth = "Invalid date";
        }
    }

    if (Object.keys(state.errors).length > 0) {
        return {
            ...prevState,
            errors: state.errors,
            error: "Please fix the highlighted fields",
        };
    }

    // here i can send data to backend
    redirect("/")
}