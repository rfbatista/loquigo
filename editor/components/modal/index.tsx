import { Paragraph } from '@rsuite/icons';
import React from 'react';
import { Modal, Button, Loader } from 'rsuite';

type SystemModalType = {
  title?: string;
  open: boolean;
  handleClose: () => void;
	handleSubmit: () => void;
	isLoading?: boolean;
};

const SystemModal: React.FC<SystemModalType> = ({
  open,
  title = "",
  handleClose,
	handleSubmit,
  children,
	isLoading = false,
}) => {
	  return (
    <Modal backdrop={true} keyboard={false} open={open} onClose={handleClose}>
			{isLoading && <Loader backdrop center content='Carregando' />}
      <Modal.Header closeButton={false}>
        <Modal.Title>{title}</Modal.Title>
      </Modal.Header>
      <Modal.Body>{children}</Modal.Body>
      <Modal.Footer>
        <Button onClick={handleSubmit} className='bg-blue-600 text-white'>
          Ok
        </Button>
        <Button onClick={handleClose} className='bg-red-600 text-white'>
          Cancelar
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default SystemModal;
