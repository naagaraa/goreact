import type { MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = () => {
  return [
    { title: "terms of service pages" },
    { name: "description", content: "terms of service pages" },
  ];
};

export default function Index() {
  return (
    <div className="font-sans p-4">
      <h1 className="text-3xl">terms of service pages</h1>
    </div>
  );
}
