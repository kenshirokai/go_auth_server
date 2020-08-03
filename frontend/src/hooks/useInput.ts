import { useState, useCallback, ChangeEvent } from 'react';

const useInput = (initialState?: string) => {
    const [state, update] = useState(initialState||'');
    const onChange = useCallback((e: ChangeEvent<HTMLInputElement>) => {
        update(e.target.value);
    }, []);
    return {state, onChange}
}

export default useInput