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

        </div>
    
    </div>


</div>

