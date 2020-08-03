import { useState, useCallback } from "react";
import { HttpCallback } from "../types/Http";

const usePost = <T>(url: string, body: any, callbacks?: HttpCallback) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [data, setData] = useState<T | null>(null);
  const [error, setError] = useState<Error | null>(null);
  const httpPost = useCallback(async () => {
    setLoading(true);
    try {
      const response = await fetch(url, {
        method: "POST",
        cache: "default",
        mode: "cors",
        headers: {
          "Content-Type": "appication/json",
        },
        body: JSON.stringify(body),
      } as RequestInit);
      setLoading(false);
      if (response.ok) {
        const responseBody = (await response.json()) as T;
        setData((_data) => ({
          ...data,
          ...responseBody,
        }));
        if (callbacks && callbacks.success) {
          callbacks.success();
        }
      } else {
        
        setError(new Error(`Network error !! status: ${response.status}`));
        if (callbacks && callbacks.failure) {
          callbacks.failure();
        }
      }
    } catch (e) {
        alert("catch")
      setError(new Error(e));
      setLoading(false);
      if (callbacks && callbacks.failure) {
        callbacks.failure();
      }
    }
  }, [url, body]);
  return {
    loading,
    data,
    error,
    httpPost,
  };
};

export default usePost;
