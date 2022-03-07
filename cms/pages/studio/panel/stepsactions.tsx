import { IconButton } from 'rsuite';
import PlusRoundIcon from '@rsuite/icons/PlusRound';
import ChatBubble from 'components/chatbuble';
import { useUpdateStepMutation } from 'services/loquiapi';

const StepActions = ({ step }) => {
	let payload = step 
  const [updateStep, { isLoading: isUpdating }] = useUpdateStepMutation({
    fixedCacheKey: payload.id,
  });
	console.log(payload)
  return (
    <>
      <div
        style={{ backgroundColor: '#F7F7FA' }}
        className='h-full z-50 p-4 max-w-sm border shadow-md sm:p-6 dark:bg-gray-800 dark:border-gray-700'
      >
        <h5 className='mb-3 text-base font-semibold text-gray-900 lg:text-xl dark:text-white'>
          Componentes
        </h5>
        {!payload.data?.components ? (
          <></>
        ) : (
          payload.data?.components?.map((component) => {
            <ChatBubble data={component} />;
          })
        )}
        <ul className='w-48 text-sm font-medium text-gray-900 bg-whitex dark:text-white'>
          <li className='w-200 my-2'>
            <IconButton
              icon={<PlusRoundIcon />}
              onClick={() => {
                const components = payload.data?.components
                  ? [
                      ...payload.data.components,
                      {
                        flowId: payload.data.flowId,
                        stepId: payload.data.id,
                        type: 'text',
                        data: { text: '' },
                        sequence: payload.components.length,
                      },
                    ]
                  : [
                      {
                        flowId: payload.data.flowId,
                        stepId: payload.data.id,
                        type: 'text',
                        data: { text: '' },
                        sequence: 0,
                      },
                    ];
                updateStep({...payload, data: {...payload.data, components}});
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
