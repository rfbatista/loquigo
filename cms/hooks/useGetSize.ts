import React from 'react';
import { useState, useCallback, useRef, MutableRefObject } from 'react';
import useEventListener from './useEventListerner';
import useIsomorphicLayoutEffect from './useIsomorphicLayoutEffect';

interface Size {
  width: number;
  height: number;
}
const useGetSize = (ref: MutableRefObject<any>): [Size] => {
  // Mutable values like 'ref.current' aren't valid dependencies
  // because mutating them doesn't re-render the component.
  // Instead, we use a state as a ref to be reactive.
  const [size, setSize] = useState<Size>({
    width: 0,
    height: 0,
  });

  // Prevent too many rendering using useCallback
  const handleSize = useCallback(() => {
    setSize({
      width: ref.current?.offsetWidth || 0,
      height: ref.current?.offsetHeight || 0,
    });

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // useEventListener('resize', handleSize);

  React.useEffect(() => {
    handleSize();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [ref.current?.offsetWidth, ref.current?.offsetHeight]);

  return [size];
};

export default useGetSize;
