import { createAsyncThunk, createEntityAdapter, createSlice } from '@reduxjs/toolkit';
import { useUpdateStepMutation } from 'services/loquiapi';
import { IStep } from '../../types/step';
import { IStepNode } from '../../types/stepnode';

const addComponentInStep = createAsyncThunk(
  'users/fetchByIdStatus',
  async (userId, thunkAPI) => {
		const [] = useUpdateStepMutation()
    const response = await userAPI.fetchById(userId)
    return response.data
  }
)

const stepAdapter = createEntityAdapter<IStepNode>({
  selectId: (stepnode) => stepnode.id,
});

const stepSlice = createSlice({
  name: 'steps',
  initialState: stepAdapter.getInitialState({
    activeStep: {},
  }),
  reducers: {
    stepAdded: stepAdapter.addOne,
    stepsAdded: stepAdapter.addMany,
    stepsReceived: stepAdapter.setAll,
		editStep: stepAdapter.updateOne,
    setActiveStep: (state, action) => {
      state.activeStep = action.payload;
    },
  },
});

const selectStepById = (id) => (state) => stepSelector.selectById(state, id);
const selectActiveStep = (state) => {
  return state?.steps?.activeStep;
};

const stepSelector = stepAdapter.getSelectors();
const { reducer, actions } = stepSlice;
const { stepAdded, stepsAdded, stepsReceived, setActiveStep } = actions;
export {
  selectActiveStep,
  selectStepById,
  stepSelector,
  stepAdded,
  stepsAdded,
  stepsReceived,
  setActiveStep,
};
export default reducer;
