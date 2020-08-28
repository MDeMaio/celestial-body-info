<template>
<div>
    <div class="row justify-content-center">
        <div class="col-md-6 list">
            <div class="input-group mb-3">
                <input type="text" class="form-control" placeholder="Search by name" v-model="title" />
                <div class="input-group-append">
                    <button class="btn btn-outline-secondary" type="button" @click="searchName">
                        Search
                    </button>
                </div>
            </div>
        </div>
    </div>
    <div class="row justify-content-center">
        <div class="col-md-6">
            <h4>Planets List</h4>
            <ul class="list-group">
                <li class="list-group-item lgi-pointer" :class="{ active: index == currentIndex }" v-for="(planet, index) in planets" :key="index" @click="setActivePlanet(planet, index)">
                    {{ planet.name }}
                </li>
            </ul>
        </div>
    </div>
    <h1 style="text-align: center; margin-top: 2%; font-size: 50px;" v-if="currentPlanet">{{ currentPlanet.name }}</h1>
    <div class="row mt-3">
        <div class="col-md-4">
            <div v-if="currentPlanet">
                <!-- <h1>{{ currentPlanet.name }}</h1> -->
                <div>
                    <label><strong>Atmosphere Info:</strong></label>
                    <ul class="list-group">
                        <li class="list-group-item"><strong>Surface Pressure: </strong>{{currentPlanet.atmosphereinfo.surfacepressure}}</li>
                        <li class="list-group-item"><strong>Elemental Composition: </strong>
                            <ul class="list-group">
                                <li class="list-group-item" v-for="(info2, index2) in currentPlanet.atmosphereinfo.element" :key="index2">
                                    <strong>{{info2.name}}:</strong> {{info2.percentasdecimal}}%
                                </li>
                            </ul>
                        </li>
                    </ul>
                </div>
            </div>
            <div v-else>

            </div>
        </div>
        <div class="col-md-4">
            <div v-if="currentPlanet">
                <!-- <h1>{{ currentPlanet.name }}</h1> -->
                <div>
                    <label><strong>Physical Info:</strong></label>
                    <ul class="list-group">
                        <li class="list-group-item" v-for="(value, key) in currentPlanet.physicalinfo" v-bind:key="key">
                            <strong>{{key}}:</strong> {{ value }}
                        </li>
                    </ul>
                </div>
            </div>
            <div v-else>

            </div>
        </div>
        <div class="col-md-4">
            <div v-if="currentPlanet">
                <div>
                    <label><strong>Orbital Info:</strong></label>
                    <ul class="list-group">
                        <li class="list-group-item" v-for="(value, key) in currentPlanet.orbitalinfo" v-bind:key="key">
                            <strong>{{key}}:</strong> {{ value }}
                        </li>
                    </ul>
                </div>
            </div>
            <div v-else>
            </div>
        </div>
    </div>
</div>
</template>

<script>
import PlanetService from "../services/PlanetService";

export default {
    name: "planets-list",
    data() {
        return {
            planets: [],
            currentPlanet: null,
            currentIndex: -1,
            title: ""
        };
    },
    methods: {
        retrievePlanets() {
            PlanetService.getAll()
                .then(response => {
                    this.planets = response.data;
                    console.log(response.data);
                })
                .catch(e => {
                    console.log(e);
                });
        },

        refreshList() {
            this.retrievePlanets();
            this.currentPlanet = null;
            this.currentIndex = -1;
        },

        setActivePlanet(planet, index) {
            this.currentPlanet = planet;
            this.currentIndex = index;
        },

        searchName() {
            PlanetService.findByTitle(this.title)
                .then(response => {
                    this.planets = response.data;
                    console.log(response.data);
                })
                .catch(e => {
                    console.log(e);
                });
        }
    },
    mounted() {
        this.retrievePlanets();
    }
};
</script>

<style>
.lgi-pointer{
  cursor: pointer;
}
.lgi-pointer:hover{
  background-color: #53a5fc;
}
</style>