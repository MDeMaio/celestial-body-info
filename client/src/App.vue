<template>
<div id="app" class="app">
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

    <div class="container-fluid mt-3">
        <transition name="fade" mode="out-in">
            <router-view />
        </transition>
    </div>

    <footer>
        <div class="container-fluid" style="background:#C8C8C8;">
            <div class="row justify-content-md-center" style="font-size:16px;padding-top:20px;">
                <div class="col-md-1" ><a href="#!">Link1</a></div>
                <div class="col-md-1"><a href="#!">Link2</a></div>
                <div class="col-md-1"><a href="#!">Link3</a></div>
                <div class="col-md-1"><a href="#!">Link4</a></div>
            </div>
            <div class="footer-copyright text-center py-3">© {{ new Date().getFullYear() }} Copyright</div>
        </div>
    </footer>
</div>
</template>

<script>
export default {
    name: 'app'
}
</script>

<style>
.app {
    font-family: Cambria, Cochin, Georgia, Times, 'Times New Roman', serif;
    font-size: 22px;
}

.footer {
    font-size: 1px;
}

html {
  overflow-y: scroll;
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
