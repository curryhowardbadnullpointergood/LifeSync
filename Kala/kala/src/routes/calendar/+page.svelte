<script>

import "./calendar.scss"
import Navbar from '../../lib/components/Navbar.svelte';
import Sidebar from '../../lib/components/Sidebar.svelte';

let isSidebarOpen = false

</script>


<div class="main">
  <Navbar on:toggleSidebar={() => isSidebarOpen = !isSidebarOpen}/>
  <Sidebar open={isSidebarOpen} />

    <div class="block">

    </div>
  
    



</div>
