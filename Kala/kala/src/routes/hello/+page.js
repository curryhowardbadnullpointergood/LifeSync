export async function load({ fetch }) {
  const res = await fetch('/api/gettask');
  const tasks = await res.json();
  return { tasks };
}// src/routes/hello/+page.js

