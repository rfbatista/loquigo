import {
  createSlice,
} from '@reduxjs/toolkit';

const botSlice = createSlice({
  name: 'bots',
  initialState: {
    activeBotId: '',
  },
  reducers: {
    setActiveBotId: (state, action) => {
      state.activeBotId = action.payload;
      return state;
    },
  },
});

const selectActiveBotId = (state) => {
  return state.bots.activeBotId;
};

export { selectActiveBotId };
export const { setActiveBotId } = botSlice.actions;
export default botSlice.reducer;
