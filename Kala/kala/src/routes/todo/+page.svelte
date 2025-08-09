<script>
import './todo.scss'

let taskInput = '';
let pointInput = '';

async function handleAdd() {
  try {
    await addTask(taskInput, pointInput);
    taskInput = '';
    pointInput = '';
  } catch (e) {
    alert('Error adding task.');
  }
}



</script>


<div class="tasks">
  
  <!-- <h1>Tasks List: </h1> --> 

  <div class="tools">
    <button type="button" >Add!</button>
    <input bind:value={taskInput} placeholder="Enter task: " />
    
    <!-- Maybe handle this as a  -->
    <input bind:value={pointInput} placeholder="Enter Points:" />
    <!--
    <button on:click={handleAdd}>Add Task</button>
    <form> 
      <input type="text" id="addTask" name="addTask" placeholder=" Enter Task: ">
    </form> -->
  </div>
 

 
  <!-- <div class="taskContainer">
    {#each data.tasks as task (task.id)}
      <div class="taskCard">
        <h3>Task {task.id}</h3>
        <p>{task.task}</p>
        <p>Points: {task.points}</p>
      </div>
    {/each}
  </div> -->



  <div class="taskContainer">
      <div class="taskCard">
        
        <!-- DUMMY TADKS  -->
        <div class="taskOverview">
          <h3>Task 1</h3>
          <button >  TASK</button>
          <p> “Nobody has ever measured, not even poets, how much the heart can hold.” ― Zelda Fitzgerald </p>
          <p>Points: Infinity</p>
        </div> 

        <div id="taskDetails1" class="taskDetails">
          Poetry is pretty, especially when you have organised time for it!
        </div>
  
      </div>

      <div class="taskCard">
        <div class="taskOverview">
        <!-- DUMMY TADKS  -->
        <h3>Task 2</h3>
        <p> “Nobody has ever measured, not even poets, how much the heart can hold.”
― Zelda Fitzgerald </p>
        <p>Points: Infinity</p>
        </div>
      </div>


      <div class="taskCard">
        <div class="taskOverview">
        <!-- DUMMY TADKS  -->
        <h3>Task 3</h3>
        <p> “Nobody has ever measured, not even poets, how much the heart can hold.”
― Zelda Fitzgerald </p>
        <p>Points: Infinity</p>
        </div>
      </div>

      <div class="taskCard"> 
        <div class="taskOverview">
        <!-- DUMMY TADKS  -->
        <h3>Task 4</h3>
        <p> “Nobody has ever measured, not even poets, how much the heart can hold.”
― Zelda Fitzgerald </p>
        <p>Points: Infinity</p>
        </div> 
      </div>




  </div>


</div>


