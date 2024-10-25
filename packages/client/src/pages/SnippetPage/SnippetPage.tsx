import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

export const SnippetPage = () => {
  const { id } = useParams();

  const [snippet, setSnippet] = useState<string>("");

  useEffect(() => {
    fetch(`http://localhost:8080/api/snippets/${id}`, {
      method: "GET",
    })
      .then((res) => res.json())
      .then((json) => {
        setSnippet(json.data);
        return;
      });
  }, []);

  return <div>{snippet}</div>;
};
