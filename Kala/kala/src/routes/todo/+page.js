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

function toggleDetails(card) {
    card.classList.toggle('expanded');
  }


//export async function addTask(task, points) {
//  try {
//    const res = await fetch('/api/addtask', {
//      method: 'POST',
//      headers: {
//        'Content-Type': 'application/json',
//      },
//      body: JSON.stringify({
//        task: task,
//        points: points
//      }),
//    });
//
//    if (!res.ok) {
//      const errText = await res.text();
//      throw new Error(`Server error ${res.status}: ${errText}`);
//    }
//
//    const data = await res.json();
//    console.log('Response from server:', data);
//    return data;
//  } catch (err) {
//    console.error('Failed to add task:', err.message);
//    throw err;
//  }
//}
//
