import React from 'react';
import Button from 'rsuite/Button';

const Panel = ({ updateBot, isLoading }) => {
  return (
    <>
      <div className='grid place-items-center'>
        <Button
          onClick={updateBot}
          loading={isLoading}
          className='bg-blue-600'
          appearance='primary'
        >
          Atualizar Bot
        </Button>
      </div>
    </>
  );
};

export default Panel;
