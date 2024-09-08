import type { MetaFunction } from "@remix-run/node";
import { Link } from "@remix-run/react";

export const meta: MetaFunction = () => {
  return [
    { title: "cms pages" },
    { name: "description", content: "cms pages" },
  ];
};

export default function Index() {
  return (
    <div className="font-sans p-4">
      <h1 className="text-3xl">cms pages</h1>
      <Link to="/cms/blog">Go to First Post</Link>
    </div>
  );
}
