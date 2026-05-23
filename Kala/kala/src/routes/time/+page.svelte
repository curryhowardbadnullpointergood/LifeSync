<script>

import "./time.scss"
import Navbar from '../../lib/components/Navbar.svelte';
import Sidebar from '../../lib/components/Sidebar.svelte';

let isSidebarOpen = false

</script>


<div class="main">
  <Navbar on:toggleSidebar={() => isSidebarOpen = !isSidebarOpen}/>
  <Sidebar open={isSidebarOpen} />

    <div class="timer">
    
        <div class="container">
            This is where you can adjust time. 
              <input type="number" id="hInput" value="0" min="0" style="width: 40px;">h
              <input type="number" id="mInput" value="5" min="0" max="59" style="width: 40px;">m
              <input type="number" id="sInput" value="0" min="0" max="59" style="width: 40px;">s
              
              <button id="startBtn">Start</button>
              
              <!-- Text Display -->
              <h1 id="display">00:05:00</h1>
        
        </div>

    
    </div>


</div>

