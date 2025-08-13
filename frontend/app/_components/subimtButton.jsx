import './components.css'
import { useFormStatus } from "react-dom";

export default function SubmitButton({normalText , pendingText}) {
    const { pending } = useFormStatus();

    return (
        <button
            type="submit"
            disabled={pending ? true : false}
            className="btn-primary"
        >
            {pending ? (
                pendingText
            ) : (
                normalText
            )}
        </button>
    );
}