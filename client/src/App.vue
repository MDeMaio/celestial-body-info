<template>
<div id="app" class="app">
    <!-- --- Application header(nav), the current routes name is used to determine if the active class is applied or not(turns the name white) ---  -->
    <nav class="navbar navbar-expand navbar-dark bg-dark">
        <a href="/" class="navbar-brand">Celestial Body Info</a>
        <div class="navbar-nav mr-auto">
            <li class="nav-item" :class="{'active':($route.name == 'planets-list')}">
                <a href="/planets" class="nav-link">Planets</a>
            </li>
            <li class="nav-item" :class="{'active':($route.name == 'stars-list')}">
                <a href="/stars" class="nav-link">Stars</a>
            </li>
        </div>
    </nav>
    <!-- --- End application header(nav) --- -->

    <!-- --- We use this for all content between the header and footer, all content has a fade in and out applied to it, 
    meaning on the first visit of each page content will fade in --- -->
    <div class="container-fluid mt-3 mb-5">
        <transition name="fade" mode="out-in">
            <router-view />
        </transition>
    </div>
    <!-- --- End application content --- -->

    <!-- --- Our footer is very basic, it is styled to be sticky to the bottom of the page and only show when scrolled all the way down --- -->
    <footer class="footer">
        <div class="container-fluid bg-dark">
            <div class="row justify-content-md-center" style="font-size:16px;padding-top:10px;">
                <div class="footer-copyright text-center py-1 text-white">© {{ new Date().getFullYear() }} Copyright</div>
                <div class="text-right py-1 text-white ml-4">
                    Maintained by <a href="https://www.linkedin.com/in/matthew-demaio-6848ab15a/" target="_blank">Matthew DeMaio</a> & <a href="https://www.linkedin.com/in/dan-pezella-495897150/" target="_blank">Dan Pezella</a>
                </div>
                <div class="text-right py-1 text-white ml-4">
                    View the code on <a href="https://github.com/MDeMaio/celestial-body-info" target="_blank">Github</a>
                </div>
            </div>
        </div>
    </footer>
    <!-- --- End application footer --- -->
</div>
</template>

<script>
export default {
    name: 'app'
}
</script>

<style>
html,
body {
    height: 100%;
}

.app {
    font-family: Cambria, Cochin, Georgia, Times, 'Times New Roman', serif;
    font-size: 22px;
    
}

.footer {
    font-size: 1px;
    position: absolute;
    bottom: 0;
    width: 100%;
}

/* We always show the scroll bar to help with a weird repositioning issue we were having, might be a better fix */
html {
    overflow-y: scroll;
    height: auto;
    position: relative;
    min-height: 100vh;
}

body{
background: #f5f5f5;

}

.fade-enter-active,
.fade-leave-active {
    transition-duration: 2s;
    transition-property: opacity;
    transition-timing-function: ease;
}

.fade-enter,
.fade-leave-active {
    opacity: 0
}
</style>
