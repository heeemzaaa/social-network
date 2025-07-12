import Button from '@/app/_components/button';
import "./modal.css";

export default function Modal({ isModalOpen, modalContent, onClose }) {
    if (isModalOpen !== true) {
        return null;
    }
    console.log('Modal received modalContent:', modalContent);
    return (
        <section className="modal">
            <article className="modal-content">
                <Button className="modal-close" onClick={onClose}> X </Button>
                {modalContent || <div>No content provided</div>} {/* Fallback for debugging */}
            </article>
        </section>
    )
}

