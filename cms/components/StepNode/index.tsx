import ReactFlow, { Handle, Position } from 'react-flow-renderer';

const customNodeStyles = {
  background: '#9CA8B3',
  color: '#FFF',
  padding: 10,
};

const StepNode = ({ data }) => {
  return (
    <div style={customNodeStyles}>
      <Handle
        type='target'
        position={Position.Left}
        style={{ borderRadius: 0 }}
				id='a'
      />
      <div>{data.text}</div>
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
