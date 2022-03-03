import React from 'react';
import { useSelector } from 'react-redux';
import { selectActiveFlow } from 'store/flow';
import Button from 'rsuite/Button';
import {
  useCreateStepMutation,
  useDeleteStepMutation,
  useGetStepQuery,
} from 'services/loquiapi';
import { IconButton, Loader } from 'rsuite';
import WarningRoundIcon from '@rsuite/icons/WarningRound';
import Actions from './panel/actions';
import { selectActiveStep, selectStepById, stepSelector } from 'store/step';
import StepActions from './panel/stepsactions';

const Panel = () => {
  const activeStep = useSelector(selectActiveStep);
  console.log(activeStep);
  return <>{activeStep?.id ? <StepActions /> : <Actions />}</>;
};

export default Panel;
