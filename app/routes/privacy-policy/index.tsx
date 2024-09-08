import type { MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = () => {
  return [
    { title: "privacy policy pages" },
    { name: "description", content: "privacy policy pages" },
  ];
};

export default function Index() {
  return (
    <div className="font-sans p-4">
      <h1 className="text-3xl">privacy policy pages</h1>
    </div>
  );
}
