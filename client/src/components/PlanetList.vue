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
                    <li :class="{'disabled':currentPage === 1}" class="page-item previous-item">
                        <router-link :to="{ query: { page: currentPage - 1 }}" class="page-link">Previous</router-link>
                    </li>
                    <li v-for="page in pageArray" :key="page" class="page-item" :class="{'active':currentPage === page}">
                        <router-link :to="{ query: { page: page }}" class="page-link">{{page}}</router-link>
                    </li>
                    <li :class="{'disabled':currentPage === totalPages}" class="page-item next-item">
                        <router-link :to="{ query: { page: currentPage + 1 }}" class="page-link">Next</router-link>
                    </li>
                </ul>
            </nav>
        </div>
    </div>
    <h1 style="text-align: center; margin-top: 2%; font-size: 50px;" v-if="currentPlanet">{{ currentPlanet.name }}</h1>
    <div class="row">
        <div class="col-md-4" v-if="currentPlanet">
            <h1>Basic Information</h1>
            <ul class="list-group">
                <li class="list-group-item">
                    <p style="font-size: 22px;"><strong>Alternate Name(s):</strong> {{currentPlanet.basic_information.alternate_name}}</p>
                </li>
                <li class="list-group-item">
                    <p style="font-size: 22px;"><strong>Number of Satelites:</strong> {{currentPlanet.basic_information.number_of_satelites > 0 ? currentPlanet.basic_information.number_of_satelites : 0}}</p>
                </li>
                <li class="list-group-item">
                    <p style="font-size: 22px;"><strong>Star System:</strong> {{currentPlanet.basic_information.star_system}}</p>
                </li>
                <li class="list-group-item">
                    <p style="font-size: 22px;"><strong>Most Abundant Resource:</strong> {{currentPlanet.basic_information.most_abundant_resource}}</p>
                </li>
                <li class="list-group-item">
                    <p style="font-size: 22px;"><strong>Surface Gravity:</strong> {{currentPlanet.basic_information.surface_gravity}}m/s<sup>2</sup></p>
                </li>
            </ul>
        </div>
        <img class="img-center col-md-4 justify-content-right" v-if="currentPlanet" v-bind:src="currentPlanet.image">
        <div class="col-md-4">
        </div>
    </div>
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
            name: "",
            pageArray: [],
            currentPage: 1,
            totalPages: 0
        };
    },
    methods: {
        retrievePlanets() {
            PlanetService.getAll(this.currentPage)
                .then(response => {
                    this.planets = response.data.planets;
                    console.log(response.data);
                    this.pageArray = this.generatePaginationPageArray(response.data.number_of_documents);
                    console.log(this.pageArray);
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
            document.getElementsByClassName("pagination")[0].style.visibility = "visible";
        },

        setActivePlanet(planet, index) {
            this.currentPlanet = planet;
            this.currentIndex = index;
        },

        searchName() {
            if (this.name === "") {
                alert("Please enter a planet to search for.");
                //this.name = "";
                this.$refs.name.focus();
                return;
            }

            PlanetService.get(this.name)
                .then(response => {
                    this.planets = [];
                    this.planets.push(response.data);
                    this.currentPlanet = null;
                    this.currentIndex = -1;
                    document.getElementsByClassName("pagination")[0].style.visibility = "hidden";
                })
                .catch(e => {
                    if (e.response.status == 404) {
                        alert("No planet by that name exists in the database.");
                        this.name = "";
                        this.$refs.name.focus();
                    }
                });
        },

        generatePaginationPageArray(numOfDocuments) {
            if (numOfDocuments == 0) { // No paging if no records.
                return;
            }

            const pages = Math.ceil(numOfDocuments / 5); // Round up for pages.
            this.totalPages = pages;
            const pageArray = [];
            for (var i = 1; i <= pages; i++) {
                pageArray.push(i);
            }

            return pageArray;
        }
    },
    watch: {
        '$route.query.page': {
            immediate: true,
            handler(page) {
                page = parseInt(page) || 1;
                if (page !== this.currentPage) {
                    PlanetService.getAll(page)
                        .then(response => {
                            this.planets = response.data.planets;
                            this.currentPage = page;
                            this.currentPlanet = null;
                            this.currentIndex = -1;
                        })
                        .catch(e => {
                            console.log(e);
                        });
                }
            }
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
