<script>
import './todo.scss'
// import {addTask } from './+page.js'
export let data;

let taskInput = '';
let pointInput = 1;

async function handleAdd() {
  try {
    await addTask(taskInput, pointInput);
    taskInput = '';
    pointInput = 1;
  } catch (e) {
    alert('Error adding task.');
  }
}
</script>



<div class="tasks">
  
  <h1>Tasks List: </h1> 

  <div class="tools">
    <button type="button" >Add!</button>
    <input bind:value={taskInput} placeholder="Enter task" />
    <input type="number" bind:value={pointInput} />
    <!--
    <button on:click={handleAdd}>Add Task</button>
    <form> 
      <input type="text" id="addTask" name="addTask" placeholder=" Enter Task: ">
    </form> -->
  </div>
 

 
  <div class="taskContainer">
    {#each data.tasks as task (task.id)}
      <div class="taskCard">
        <h3>Task {task.id}</h3>
        <p>{task.task}</p>
        <p>Points: {task.points}</p>
      </div>
    {/each}
  </div>


</div>


