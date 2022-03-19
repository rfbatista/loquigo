import { IconButton } from 'rsuite';
import PlusRoundIcon from '@rsuite/icons/PlusRound';
import ChatBubble from 'components/chatbuble';
import {
  useDeleteComponentMutation,
  useGetStepByIdQuery,
  useUpdateStepMutation,
} from 'services/loquiapi';
import React from 'react';

const addComponent = (step, newComponent) => {
  return step?.components?.length > 0
    ? [
        ...step.components,
        {
          flowId: step.flowId,
          stepId: step.id,
          type: newComponent.type,
          sequence: step.components?.length,
        },
      ]
    : [
        {
          flowId: step.flowId,
          stepId: step.id,

          sequence: 0,
        },
      ];
};

const StepActions = ({ step }) => {
  const { data, isError, isLoading, error } = useGetStepByIdQuery(step.id);
  const [updateStep, { isLoading: isUpdating, isSuccess }] =
    useUpdateStepMutation({
      fixedCacheKey: step.id,
    });
  const [deleteStep] = useDeleteComponentMutation();
  const addComponentInStep = (component) => {
    updateStep({
      ...data,
      data: { ...data.data, components: addComponent(data, component) },
    });
  };
  const removeComponent = (component) => {
    deleteStep(component);
  };
  const list = data?.Components?.map((component, index) => (
    <ChatBubble
      key={index.toString()}
      data={component}
      remove={removeComponent}
    />
  ));

  return (
    <>
      <div
        style={{ backgroundColor: '#F7F7FA' }}
        className='h-full z-50 p-4 max-w-sm border shadow-md sm:p-6 dark:bg-gray-800 dark:border-gray-700'
      >
        <h5 className='mb-3 text-base font-semibold text-gray-900 lg:text-xl dark:text-white'>
          Componentes
        </h5>
        {list}
        <ul className='w-48 text-sm font-medium text-gray-900 bg-whitex dark:text-white'>
          <li className='w-200 my-2'>
            <IconButton
              icon={<PlusRoundIcon />}
              onClick={() => {
                addComponentInStep({
                  key: new Date().getTime(),
                  type: 'text',
                  data: { text: '' },
                });
              }}
            >
              Adicionar Texto
            </IconButton>
          </li>
          <li className='w-200 my-2'>
            <IconButton icon={<PlusRoundIcon />}>Adicionar Hold</IconButton>
          </li>
        </ul>
      </div>
    </>
  );
};

export default StepActions;
