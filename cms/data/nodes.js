export const elements = [
	{
		id: '1',
		type: 'step', // input node
		data: { text: 'Input Node' },
		position: { x: 250, y: 25 },
	},
	// default node
	{
		id: '2',
		// you can also pass a React component as a label
		data: { label: 'Default' },
		position: { x: 100, y: 125 },
	},
	{
		id: '3',
		type: 'output', // output node
		data: { label: 'Output Node' },
		position: { x: 250, y: 250 },
	},
	{
		id: '4',
		type: 'output', // output node
		data: { label: 'Output Node' },
		position: { x: 250, y: 250 },
	},
	// animated edge
	{ id: 'e1-2', source: '1', target: '2', animated: true, targetHandle: 'a' },
	{ id: 'e2-3', source: '2', target: '3' },
	{ id: 'e2-4', source: '1', target: '3', sourceHandle: 'b' },
	{ id: 'e2-5', source: '1', target: '4', sourceHandle: 'c' },
];