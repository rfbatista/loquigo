import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import config from 'config';
import mockBaseQuery from './mockBaseQuery';
import axiosBaseQuery from './axiosBaseQuery';
import { IStep } from 'types/step';

export const FLOW_API_REDUCER_KEY = 'FLOW_API';
export const loquiapi = createApi({
  reducerPath: FLOW_API_REDUCER_KEY,
  tagTypes: ['Step'],
  baseQuery: axiosBaseQuery({
    baseUrl: String(config.core.endpoint),
  }),
  endpoints: (builder) => ({
    getFlow: builder.query({
      query: (botId) => ({ url: `/flow/${botId}`, method: 'GET', data: null }),
    }),
    getStep: builder.query<IStep[], void>({
      query: (flowId) => ({
        url: `/step/flow/${flowId}`,
        method: 'GET',
        data: null,
      }),
      providesTags: (result, error, arg) =>
        result
          ? [
              ...result.map(({ flow_id }) => ({
                type: 'Step' as const,
                id: flow_id,
              })),
              'Step',
            ]
          : ['Step'],
    }),
    createStep: builder.mutation({
      query: (step) => ({ url: `/step/`, method: 'POST', data: step }),
      invalidatesTags: ['Step'],
    }),
    updateStep: builder.mutation({
      query: (step) => ({ url: `/step/${step.id}`, method: 'PUT', data: step }),
      invalidatesTags: ['Step'],
    }),
    deleteStep: builder.mutation({
      query: (step) => ({ url: `/step/`, method: 'DELETE', data: step }),
      invalidatesTags: ['Step'],
    }),
    getFlowMap: builder.query<IStepNode[], void>({
      query: (flowId) => ({
        url: `/flow/map/${flowId}`,
        method: 'GET',
        data: null,
      }),
      providesTags: (result, error, arg) =>
        result
          ? [
              ...result.map(({ id }) => ({ type: 'Step' as const, id: id })),
              'Step',
            ]
          : ['Step'],
    }),
  }),
});

export const {
  useGetFlowQuery,
  useGetStepQuery,
  useGetFlowMapQuery,
  useCreateStepMutation,
  useDeleteStepMutation,
	useUpdateStepMutation,
} = loquiapi;
