import { useEffect, useState } from 'react';

const useFetch = <T>(url: string, options:RequestInit) => {
    
    const [data, setData] = useState<T|null>(null);
    const [error, setError] = useState<Error|null>(null);
    const [loading, setLoading] = useState<boolean>(false);

    useEffect(() => {
        const abort = new AbortController();
        (async () => {
            setLoading(true)
            const data = await fetch(url, {...options, signal: abort.signal})
            if(!data.ok) {
                setError(new Error(`status:${data.status}`))
            }else {
                const responseBody = await data.json() as T
                setData(responseBody)
                setLoading(false);
            }
        })()
        //cleanup
        return () => { abort.abort() };
    }, [url, options]);
    return { data, error, loading }
}

export default useFetch;