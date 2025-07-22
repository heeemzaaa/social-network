// app/_context/ModalContext.jsx or .tsx
'use client';

import React, { createContext, useContext, useState } from 'react';
import Modal from '../_components/modal/modal';

const ModalContext = createContext();
export const useModal = () => {
    const context = useContext(ModalContext);
    if (!context) {
        throw new Error('useModal must be used within a ModalProvider');
    }
    return context;
};

export function ModalProvider({ children }) {
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [modalContent, setModalContent] = useState(null);

    // data passed when a form is submitted successfuly.
    const [data, setData] = useState()

    const openModal = (content) => { 
        console.log("opening modal: ", content)
        setIsModalOpen(true);
        setModalContent(content);
    };

    const closeModal = () => {
        setIsModalOpen(false);
        setModalContent(null);
    };

    const setModalData = (data) => {
        setData(data)
    }

    const getModalData = () => {
        return data
    }

    return (
        <ModalContext.Provider value={{ openModal, closeModal, setModalData, getModalData }}>
            {children}
            <Modal isModalOpen={isModalOpen} modalContent={modalContent} onClose={closeModal} />
        </ModalContext.Provider>
    );
}
