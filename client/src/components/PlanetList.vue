<template>
<div>
    <div class="row justify-content-center">
        <div class="col-md-6 list">
            <form v-on:submit.prevent @submit="searchName">
                <div class="input-group mb-3">
                    <input ref="name" type="text" class="form-control fs-20" placeholder="Search by name" v-model="name" />
                    <div class="input-group-append">
                        <button class="btn btn-outline-primary ml-1 fs-20" type="submit">
                            Search
                        </button>
                        <button class="btn btn-outline-secondary fs-20" type="button" @click="refreshList">
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
            <div class="btn-group dropleft mb-2">
                <button type="button" class="btn btn-secondary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                    Planet Type
                </button>
                <div class="dropdown-menu">
                    <a class="dropdown-item" v-bind:href="'/planets?page=1&type=All'" :class="{'active':(type === 'All')}">All</a>
                    <a class="dropdown-item" v-bind:href="'/planets?page=1&type=Inner Planet'" :class="{'active':(type === 'Inner Planet')}">Inner Planet</a>
                    <a class="dropdown-item" v-bind:href="'/planets?page=1&type=Outer Planet'" :class="{'active':(type === 'Outer Planet')}">Outer Planet</a>
                    <a class="dropdown-item" v-bind:href="'/planets?page=1&type=Exoplanet'" :class="{'active':(type === 'Exoplanet')}">Exoplanet</a>
                </div>
            </div>
            <transition-group name="slide-fade" tag="ul" class="list-group">
                <li class="list-group-item lgi-pointer py-1" :class="{ active: index == currentIndex }" v-for="(planet, index) in planets" :key="planet.planet_id" @click="setActivePlanet(planet, index)">
                    {{ planet.name }}
                </li>
            </transition-group>
            <nav aria-label="Pagination">
                <transition-group name="slide-fade" tag="ul" class="pagination mt-2 justify-content-center">
                    <li v-if="loaded" :key="'prev'" :class="{'disabled':currentPage === 1}" class="page-item previous-item">
                        <a class="page-link" v-bind:href="'/planets?page='+ (currentPage-1) + '&type=' + type">Prev</a>
                    </li>
                    <li v-for="page in pageArray" :key="page" class="page-item" :class="{'active':(currentPage === page)}">
                        <a class="page-link" v-bind:href="'/planets?page='+ page + '&type=' + type">{{page}}</a>
                    </li>
                    <li v-if="loaded" :key="'next'" :class="{'disabled':currentPage === totalPages}" class="page-item next-item">
                        <a class="page-link" v-bind:href="'/planets?page='+ (currentPage+1) + '&type=' + type">Next</a>
                    </li>
                </transition-group>
            </nav>
        </div>
    </div>
    <h1 style="text-align: center; margin-top: 2%; font-size: 50px;" v-if="currentPlanet">{{ currentPlanet.name }}</h1>
    <div class="row">
        <div class="col-md-4" v-if="currentPlanet">
            <h1>Basic Information</h1>
            <ul class="list-group">
                <li :key="'alternate_name'" class="list-group-item lgi-colored-space">
                    <p style="font-size: 22px;"><strong>Alternate Name(s):</strong> {{currentPlanet.basic_information.alternate_name}}</p>
                </li>
                <li :key="'type'" class="list-group-item lgi-colored-space">
                    <p style="font-size: 22px;"><strong>Planetary Type:</strong> {{currentPlanet.basic_information.type}}</p>
                </li>
                <li :key="'number_of_satelites'" class="list-group-item lgi-colored-space">
                    <p style="font-size: 22px;"><strong>Number of Satelites:</strong> {{currentPlanet.basic_information.number_of_satelites > 0 ? currentPlanet.basic_information.number_of_satelites : 0}}</p>
                </li>
                <li :key="'star_system'" class="list-group-item lgi-colored-space">
                    <p style="font-size: 22px;"><strong>Star System:</strong> {{currentPlanet.basic_information.star_system}}</p>
                </li>
                <li :key="'most_abundant_resource'" class="list-group-item lgi-colored-space">
                    <p style="font-size: 22px;"><strong>Most Abundant Resource:</strong> {{currentPlanet.basic_information.most_abundant_resource}}</p>
                </li>
                <li :key="'surface_gravity'" class="list-group-item lgi-colored-space">
                    <p style="font-size: 22px;"><strong>Surface Gravity:</strong> {{currentPlanet.basic_information.surface_gravity}}m/s<sup>2</sup></p>
                </li>
            </ul>
        </div>
        <img class="img-center col-md-4 justify-content-right" v-if="currentPlanet" v-bind:src="currentPlanet.image">
        <div class="col-md-4">
        </div>
    </div>
    <div class="row mt-3 justify-content-center">
        <div class="col-md-10">
            <div v-if="currentPlanet">
                <ul class="list-group list-group-flush">
                    <li class="list-group-item mt-2" v-for="(val, index) in currentPlanet.facts" :key="index">
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
            currentPage: parseInt(this.$route.query.page),
            totalPages: 0,
            type: this.$route.query.type,
            loaded: false
        };
    },
    methods: {
        retrievePlanets(page, type) { // Fetchs all of our planets for the current page.
            PlanetService.getAll(page, type)
                .then(response => {
                    this.planets = response.data.planets;
                    console.log(response.data);
                    console.log(this.pageArray);
                    this.currentPage = page;
                    this.pageArray = this.generatePaginationPageArray(response.data.number_of_documents);
                    this.currentPlanet = null;
                    this.currentIndex = -1;
                })
                .catch(e => {
                    console.log(e);
                });
        },

        refreshList() { // Refreshes the page to the default state.
            this.retrievePlanets(this.currentPage, this.type);
            this.currentPlanet = null;
            this.currentIndex = -1;
            this.name = "";
            document.getElementsByClassName("pagination")[0].style.visibility = "visible";
        },

        setActivePlanet(planet, index) { // Updates currently viewed planet.
            this.currentPlanet = planet;
            this.currentIndex = index;
        },

        searchName() { // This will search for a planet by name and return it, caps specific as of now.
            if (this.name === "") {
                alert("Please enter a planet to search for.");
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

        generatePaginationPageArray(numOfDocuments) { // Fetch our pages in an array so we can iterate over it later to create the pagination list items.
            if (numOfDocuments == 0) { // No paging if no records.
                return;
            }

            const pages = Math.ceil(numOfDocuments / 5); // Round up for pages.
            this.totalPages = pages;
            const pageArray = [];

            console.log("Total Pages:" + this.totalPages + " Current Page:" + this.currentPage);
            for (let i = 1; i <= pages; i++) {
                if (this.totalPages <= 5) { // Show all if 5 or less.
                    pageArray.push(i);
                } else {
                    if (this.currentPage == 1 || this.currentPage == 2) {
                        if (i <= 5) {
                            pageArray.push(i);
                        }
                    } else {
                        if (i >= this.currentPage - 2 && i <= this.currentPage + 2) {
                            pageArray.push(i);
                        }
                    }
                }
            }

            return pageArray;
        }
    },
    watch: { // Watch for data change in which page the user is currently on, call API to get new data when it changes.
        // '$route.query.page': {
        //     immediate: true,
        //     handler(page) {
        //         page = parseInt(page) || 1;
        //         if (page !== this.currentPage) {
        //             this.retrievePlanets(page, this.type);
        //         }
        //     }
        // },
        // "currentPage": function () {
        //     this.refreshList();
        // },

        // "type": function () {
        //      if (this.currentPage === 1) {   // Otherwise changing current page will take care of the refresh for us.
        //         this.refreshList();
        //     }
        //     this.currentPage = 1;
        // }
    },
    mounted() {
        this.retrievePlanets(this.currentPage, this.type);
        this.loaded = true;
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

.lgi-colored-space {
    background-color: #fcfbfe;
    color: #1d1135;
}

.img-center {
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 50%;
}

.fs-20 {
    font-size: 20px;
}

/* Enter and leave animations can use different */
/* durations and timing functions.              */
.slide-fade-enter-active {
    transition: all 2s ease;
}

.slide-fade-leave-active {
    transition: all 1.5s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}

.slide-fade-enter,
.slide-fade-leave-to

/* .slide-fade-leave-active below version 2.1.8 */
    {
    transform: translateX(40px);
    opacity: 0;
}
</style>
