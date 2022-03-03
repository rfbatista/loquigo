export type IComponent = {
  id: string;
};
export type IStep = {
  id: string;
  flow_id: string;
  name: string;
  components: IComponent[];
};
