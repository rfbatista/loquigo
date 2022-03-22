// export type IFlow = [Step | Edge];

import { IStep } from './step';

export type IStepNode = {
  id: string;
  type?: 'output' | 'input' | 'step';
  data: IStep;
  position: {
    x: number;
    y: number;
  };
};

export type Edge = {
  id: string;
  source: string;
  target: string;
  animated: boolean;
};
