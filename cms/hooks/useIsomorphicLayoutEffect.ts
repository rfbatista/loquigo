import React from 'react';

//avoid server rendering error
const useIsomorphicLayoutEffect =
  typeof window !== 'undefined' ? React.useLayoutEffect : React.useEffect;

export default useIsomorphicLayoutEffect;
