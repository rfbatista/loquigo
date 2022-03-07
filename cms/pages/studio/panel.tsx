import React from 'react';
import { useSelector } from 'react-redux';
import Actions from './panel/actions';
import { selectActiveStep, selectStepById, stepSelector } from 'store/step';
import StepActions from './panel/stepsactions';

const Panel = () => {
  const activeStep = useSelector(selectActiveStep);
  return <>{activeStep?.id ? <StepActions step={activeStep} /> : <Actions />}</>;
};

export default Panel;
