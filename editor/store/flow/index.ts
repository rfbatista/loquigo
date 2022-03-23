import {
  createSlice,
} from '@reduxjs/toolkit';

const flowSlice = createSlice({
  name: 'flows',
  initialState: {
    activeFlow: {},
		activeBotYaml: '',
  },
  reducers: {
    setActiveFlow: (state, action) => {
      state.activeFlow = action.payload;
      return state;
    },
		updateBotYaml: (state, action) => {
			state.activeBotYaml = action.payload
			return state
		}
  },
});

const selectActiveFlow = (state) => {
  return state.flows.activeFlow;
};

const getActiveBotYaml = (state) => {
  return state.activeBotYaml;
}

export { selectActiveFlow, getActiveBotYaml };
export const { setActiveFlow, updateBotYaml } = flowSlice.actions;
export default flowSlice.reducer;
