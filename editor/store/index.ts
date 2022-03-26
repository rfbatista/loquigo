import { configureStore } from '@reduxjs/toolkit';
import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux';
import {
  FLUSH,
  PAUSE,
  PERSIST,
  persistStore,
  PURGE,
  REGISTER,
  REHYDRATE,
} from 'redux-persist';
import { setupListeners } from '@reduxjs/toolkit/query';
import { loquiapi } from 'services/loquiapi';
import flowReducer from './flow';
import stepReducer from './step';
import botsReducer from './bot';

const reducers = {
  [loquiapi.reducerPath]: loquiapi.reducer,
  flows: flowReducer,
  steps: stepReducer,
	bots: botsReducer,
};

// const combinedReducer = combineReducers<typeof reducers>(reducers);

// export const rootReducer = (state, action) => {
//   if (action.type === RESET_STATE_ACTION_TYPE) {
//     state = {} as RootState;
//   }

//   return combinedReducer(state, action);
// };

// const persistConfig = {
//   key: 'root',
//   storage,
// };

// const persistedReducer = persistReducer(persistConfig, combinedReducer);

const store = configureStore({
  reducer: reducers,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
      },
    }).concat(loquiapi.middleware),
});

export const persistor = persistStore(store);
setupListeners(store.dispatch); // NOTE this addition

export default store;

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export const useTypedDispatch = () => useDispatch<AppDispatch>();
export const useTypedSelector: TypedUseSelectorHook<RootState> = useSelector;
