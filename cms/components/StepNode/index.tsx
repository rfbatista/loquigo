import ReactFlow, { Handle, Position } from 'react-flow-renderer';
import { IconButton } from 'rsuite';
import PlusRoundIcon from '@rsuite/icons/PlusRound';
import { useDispatch } from 'react-redux';
import { setActiveStep } from 'store/step';
const customNodeStyles = {
  background: '#9CA8B3',
  color: '#FFF',
  padding: 10,
};

const StepNode = (step) => {
	const dispatch = useDispatch()
  return (
    <div
      style={{ width: '200px' }}
      className='z-50 w-200'
      onClick={() => {
				console.log('foi')
        dispatch(setActiveStep(step));
      }}
    >
      <div className='leading-7 text-sm w-200 px-2 rounded-t-lg font-bold text-black'>
        {step?.data?.name}
      </div>
      <ul className='w-48 text-sm font-medium text-gray-900 bg-whitex dark:text-white'>
        <li className='w-200 '>
          <IconButton icon={<PlusRoundIcon />}>Adicionar componente</IconButton>
        </li>
      </ul>
      <div></div>
      <Handle
        type='target'
        position={Position.Left}
        style={{ borderRadius: 0 }}
        id='a'
      />
      <Handle
        type='source'
        position={Position.Right}
        id='b'
        style={{ top: '30%', borderRadius: 0 }}
      />
      <Handle
        type='source'
        position={Position.Right}
        id='c'
        style={{ top: '70%', borderRadius: 0 }}
      />
    </div>
  );
};

export default StepNode;
