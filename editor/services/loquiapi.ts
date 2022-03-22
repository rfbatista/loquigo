import { createApi } from '@reduxjs/toolkit/query/react';
import config from 'config';
import axiosBaseQuery from './axiosBaseQuery';
import { IStep } from 'types/step';

export const FLOW_API_REDUCER_KEY = 'FLOW_API';
export const loquiapi = createApi({
  reducerPath: FLOW_API_REDUCER_KEY,
  tagTypes: ['Step', 'Bot'],
  baseQuery: axiosBaseQuery({
    baseUrl: String(config.core.endpoint),
  }),
  endpoints: (builder) => ({
    updateBot: builder.mutation({
      query: (step) => ({
        url: `/editor/`,
        method: 'PUT',
        data: { data: step },
      }),
      invalidatesTags: ['Bot'],
    }),
    getFlow: builder.query({
      query: (botId) => ({ url: `/flow/${botId}`, method: 'GET', data: null }),
    }),
    getBot: builder.query({
      query: (botId) => ({
        url: `/editor/${botId}`,
        method: 'GET',
        data: null,
      }),
    }),
    getBotVersions: builder.query({
      query: (botId) => ({
        url: `/editor/${botId}/version`,
        method: 'GET',
        data: null,
      }),
    }),
    getBotVersion: builder.query({
      query: ({ botId, version }) => ({
        url: `/editor/${botId}/version/${version}`,
        method: 'GET',
        data: null,
      }),
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
    deleteComponent: builder.mutation({
      query: (step) => ({ url: `/component/`, method: 'DELETE', data: step }),
      invalidatesTags: ['Step'],
    }),
    getStepById: builder.query({
      query: (stepId) => ({
        url: `/step/${stepId}`,
        method: 'GET',
        data: null,
      }),
      providesTags: (result, error, arg) =>
        result ? [{ type: 'Step' as const, id: result.id }, 'Step'] : ['Step'],
    }),
    getFlowMap: builder.query({
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
  useGetStepByIdQuery,
  useGetFlowMapQuery,
  useDeleteComponentMutation,
  useCreateStepMutation,
  useDeleteStepMutation,
  useUpdateStepMutation,
  useUpdateBotMutation,
  useGetBotQuery,
	useGetBotVersionQuery,
	useGetBotVersionsQuery,
} = loquiapi;
