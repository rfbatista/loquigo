import React from 'react';
import { useSelector } from 'react-redux';
import Actions from './panel/actions';
import { selectActiveStep, selectStepById, stepSelector } from 'store/step';
import StepActions from './panel/stepsactions';
import Button from 'rsuite/Button';
import { Sidebar, Sidenav } from 'rsuite';

const Panel = ({ updateBot }) => {
  return (
    <>
      <div className='grid place-items-center'>
        <Button onClick={updateBot} className='bg-blue-600' appearance='primary'>
          Atualizar Bot
        </Button>
      </div>
    </>
  );
};

export default Panel;
