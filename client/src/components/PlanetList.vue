<template>
<div>
    <div class="row justify-content-center">
        <div class="col-md-6 list">
            <form v-on:submit.prevent @submit="searchName">
            <div class="input-group mb-3">
                <input ref="name" type="text" class="form-control" placeholder="Search by name" v-model="name" />
                <div class="input-group-append">
                    <button class="btn btn-outline-primary ml-1" type="submit">
                        Search
                    </button>
                    <button class="btn btn-outline-secondary" type="button" @click="refreshList">
                        Reset
                    </button>
                </div>
            </div>
            </form>
        </div>
    </div>
    <div class="row justify-content-center">
        <div class="col-md-6">
            <h4 style="font-size: 35px;">Available Planets</h4>
            <ul class="list-group">
                <li class="list-group-item lgi-pointer py-1" :class="{ active: index == currentIndex }" v-for="(planet, index) in planets" :key="planet.planet_id" @click="setActivePlanet(planet, index)">
                    {{ planet.name }}
                </li>
            </ul>
            <nav aria-label="Page navigation example">
                <ul class="pagination mt-2 justify-content-center">
                    <li class="page-item"><a class="page-link" href="#">Previous</a></li>
                    <li class="page-item"><a class="page-link" href="#">1</a></li>
                    <li class="page-item"><a class="page-link" href="#">2</a></li>
                    <li class="page-item"><a class="page-link" href="#">3</a></li>
                    <li class="page-item"><a class="page-link" href="#">Next</a></li>
                </ul>
            </nav>
        </div>
    </div>
    <h1 style="text-align: center; margin-top: 2%; font-size: 50px;" v-if="currentPlanet">{{ currentPlanet.name }}</h1>
    <img class="img-center" v-if="currentPlanet" v-bind:src="currentPlanet.image">
    <div class="row mt-3 justify-content-center">
        <div class="col-md-8">
            <div v-if="currentPlanet">
                <ul class="list-group">
                    <li class="list-group-item" v-for="(val, index) in currentPlanet.facts" :key="index">
                        <h1 style="font-size: 40px; text-align: center;">{{val.title}}</h1>
                        <p style="font-size: 22px;">{{val.fact}}</p>
                    </li>
                </ul>
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
            name: ""
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
            this.name = "";
        },

        setActivePlanet(planet, index) {
            this.currentPlanet = planet;
            this.currentIndex = index;
        },

        searchName() {
            let id = "";
            this.planets.forEach((element) => {
                if (element.name.toLowerCase() == this.name.toLowerCase()) {
                    id = element.planet_id;
                }
            });

            if (id === "") {
                alert("No planet exists by that name.");
                this.name = "";
                this.$refs.name.focus();
                return;
            }

            PlanetService.get(id)
                .then(response => {
                    this.planets = [];
                    this.planets.push(response.data);
                    this.currentPlanet = null;
                    this.currentIndex = -1;
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
.lgi-pointer {
    cursor: pointer;
}

.lgi-pointer:hover {
    background-color: #53a5fc;
}

.img-center {
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 50%;
}
</style>
