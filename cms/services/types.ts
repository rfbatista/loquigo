// export type IFlow = [Step | Edge];

type Step = {
  id: string;
  type?: 'output' | 'input';
  data: any;
  position: {
    x: number;
    y: number;
  };
};

type Edge = {
  id: string;
  source: string;
  target: string;
  animated: boolean;
};
