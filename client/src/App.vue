<template>
<div id="app" class="app">
    <nav class="navbar navbar-expand navbar-dark bg-dark">
        <a href="/" class="navbar-brand">Celestial Body Info</a>
        <div class="navbar-nav mr-auto">
            <li class="nav-item" :class="{'active':($route.name == 'planets-list')}">
                <a href="/planets?page=1&type=All" class="nav-link">Planets</a>
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

.fade-enter-active,
.fade-leave-active {
    transition-duration: 1.5s;
    transition-property: opacity;
    transition-timing-function: ease;
}

.fade-enter,
.fade-leave-active {
    opacity: 0
}
</style>
