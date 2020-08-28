<template>
<div id="app">
    <nav class="navbar navbar-expand navbar-dark bg-dark">
        <a href="#" class="navbar-brand">Celestial Body Info</a>
        <div class="navbar-nav mr-auto">
            <li class="nav-item">
                <a href="/planets" class="nav-link">Planets</a>
            </li>
            <li class="nav-item">
                <a href="/stars" class="nav-link">Stars</a>
            </li>
        </div>
    </nav>

    <div class="container-fluid mt-3">
        <router-view />
    </div>
</div>
</template>

<script>
export default {
    name: 'app'
}
</script>

<style>
.container {
    margin-left: 100px;
}
</style>
