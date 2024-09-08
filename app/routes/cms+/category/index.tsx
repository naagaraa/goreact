import type { MetaFunction } from "@remix-run/node";
import { useEffect, useState } from "react";

export const meta: MetaFunction = () => {
  return [
    { title: "category pages" },
    { name: "description", content: "category pages" },
  ];
};

export default function Index() {
  const [message, setMessage] = useState("");
  const apiUrl = import.meta.env.VITE_API_URL;

  useEffect(() => {
    fetch(`${apiUrl}/api/v1/cms/category`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      mode: "cors", // Ensure CORS mode is enabled
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => setMessage(data.message))
      .catch((error) => console.error("Error fetching API:", error));
  }, [apiUrl]);

  return (
    <div className="font-sans p-4">
      <h1 className="text-3xl">category pages</h1>
      <div>{message}</div>
    </div>
  );
}
