// app/error.js
"use client";

import Button from "../_components/button";

// Error components must be client components

export default function Error({ error, reset }) {
    return (
        <main className="p-4">
            <br />
            <h2 >Something went wrong!</h2>
            <br />
            <p>{JSON.stringify(error)}</p>
            <br />
            <p>{error.message}</p>
            <p>{error.status}</p>
            <br />
            <button
                className="btn-danger"
                onClick={() => reset()}
            >
                Try again
            </button>
        </main>
    );
}