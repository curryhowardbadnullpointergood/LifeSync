export async function load({ fetch }) {
  const res = await fetch('/api/gettask');
  
  if (!res.ok) {
    console.error('Failed to fetch tasks:', res.status, res.statusText);
    return { tasks: [] };
  }

  const tasks = await res.json();
  console.log('Fetched tasks:', tasks); // Check the output in the browser console
  return { tasks };
}

