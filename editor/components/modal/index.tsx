import { Paragraph } from '@rsuite/icons';
import React from 'react';
import { Modal, Button } from 'rsuite';

type SystemModalType = {
  title?: string;
  open: boolean;
  handleClose: () => void;
};

const SystemModal: React.FC<SystemModalType> = ({
  open,
  title = "",
  handleClose,
  children,
}) => {
  return (
    <Modal backdrop={true} keyboard={false} open={open} onClose={handleClose}>
      <Modal.Header closeButton={false}>
        <Modal.Title>{title}</Modal.Title>
      </Modal.Header>
      <Modal.Body>{children}</Modal.Body>
      <Modal.Footer>
        <Button onClick={handleClose} className='bg-blue-600 text-white'>
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
