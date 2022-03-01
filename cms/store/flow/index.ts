import {
  createSlice,
} from '@reduxjs/toolkit';

const flowSlice = createSlice({
  name: 'flows',
  initialState: {
    activeFlow: {},
  },
  reducers: {
    setActiveFlow: (state, action) => {
      state.activeFlow = action.payload;
      return state;
    },
  },
});

const selectActiveFlow = (state) => {
  return state.flows.activeFlow;
};

export { selectActiveFlow };
export const { setActiveFlow } = flowSlice.actions;
export default flowSlice.reducer;
