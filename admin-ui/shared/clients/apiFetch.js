import { API_BASE_URL } from "./../constants";

const useApiFetch = (url, options) => {
  const [response, setResponse] = React.useState(null);
  const [error, setError] = React.useState(null);
  const [isLoading, setIsLoading] = React.useState(false);
  React.useEffect(() => {
    const fetchData = async () => {
      setIsLoading(true);
      try {
        const res = await fetch(API_BASE_URL + url, options);
        const json = await res.json();
        setResponse(json);
        setIsLoading(false);
      } catch (error) {
        setError(error);
      }
    };
    fetchData();
  }, []);
  return { response, error, isLoading };
};

const apiFetch = async () => {
  setIsLoading(true);
  try {
    const res = await fetch(API_BASE_URL + url, options);
    const json = await res.json();
    setResponse(json);
    setIsLoading(false);
  } catch (error) {
    setError(error);
  }
};
