<template>
<div>
    <search-bar @search="searchName" @reset="resetPage"> </search-bar>
    <div class="row justify-content-center">
        <div class="col-md-6">
            <h4 style="font-size: 35px;">Available Planets</h4>
            <!-- --- Filter options for planet, including search bar ---  -->
            <div class="btn-group dropleft mb-2">
                <button type="button" class="btn btn-secondary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                    Planet Type
                </button>
                <div class="dropdown-menu">
                    <button class="dropdown-item" v-for="(planetType) in planetTypes" :key="planetType" @click="type = planetType" :class="{'active':(type === planetType)}">{{planetType}}</button>
                </div>
            </div>
            <!-- --- End filter options for planet --- -->
            <!-- --- List of all our planets for current criteria --- -->
            <ul class="list-group">
                <li class="list-group-item lgi-pointer py-1" :class="{ active: index == currentIndex }" v-for="(planet, index) in planets" :key="planet.planet_id" @click="setActivePlanet(planet, index)">
                    {{ planet.name }}
                </li>
            </ul>
            <!-- --- End list of all our planets ---  -->
            <!-- --- Pagination for our planets list --- -->
            <pagination :numOfDocuments="documents" :onPageChange="handlePageChange" ref="pagination"> </pagination>
            <!-- --- End pagination for our planets list --- -->
        </div>
    </div>
    <transition name="slide-fade">
        <h1 style="text-align: center; margin-top: 2%; font-size: 50px;" v-if="currentPlanet">{{ currentPlanet.name }}</h1>
    </transition>
    <!-- --- Current planets image and basic information... we can probably break this into its own component --- -->
    <div class="row">
        <transition name="slide-fade">
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
        </transition>
        <transition name="slide-fade">
            <img class="img-center col-md-4 justify-content-right" v-if="currentPlanet" v-bind:src="'data:image/jpeg;base64,' + currentPlanet.image">
        </transition>
        <div class="col-md-4">
        </div>
    </div>
    <!-- --- End Current planet image and basic information --- -->
    <!-- --- List of our planet facts --- -->
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
    <!-- --- End list of planet facts --- -->
</div>
</template>

<script>
import PlanetService from "../services/PlanetService";

export default {
    name: "planets-list",
    data() {
        return {
            planets: [],
            planetTypes: [],
            currentPlanet: null,
            currentIndex: -1,
            documents: 0,
            type: typeof this.$route.query.type == "undefined" ? "All" : this.$route.query.type,
            loaded: false,
            resetting: false,
            searching: false
        };
    },
    methods: {
        retrievePlanets(page, type, name) { // Fetchs all of our planets for the current page.
            PlanetService.getAll(page, type, name == "" ? "All" : name)
                .then(response => {
                    this.planets = response.data.planets;
                    this.type = type;
                    this.documents = response.data.number_of_documents;
                    this.clearPlanetView();
                })
                .catch(e => {
                    console.log(e);
                }).finally(() => {
                    this.resetting = false; // why here????
                    this.searching = false; // ???????????
                });
        },

        retrievePlanetTypes() { // Fetchs all of our planet types.
            PlanetService.getPlanetTypes()
                .then(response => {
                    response.data.planet_type.unshift("All"); // Adds "All" to the front of the array, so it is the first option.
                    this.planetTypes = response.data.planet_type;
                })
                .catch(e => {
                    console.log(e);
                })
        },

        clearPlanetView() { // Clears the planet information view.
            this.currentIndex = -1;
            this.currentPlanet = null;
        },

        setActivePlanet(planet, index) { // Updates currently viewed planet.
            this.currentPlanet = planet;
            this.currentIndex = index;
        },

        searchName(value) { // This will search for a planet by name and return it, based on REGEX of value in search input box(matches anything containing that value)
            this.searching = true;
            this.retrievePlanets(1, "All", value);
            this.$router.push({
                query: Object.assign({}, this.$route.query, {
                    page: 1,
                    type: "All",
                    name: value
                })
            }).catch(() => {}); // Catch the error because we dont care about redirecting to same page.
            
            if (this.$refs.pagination.currentPage !== 1) { // Access childs current page, changing this to 1 will trigger a data reload.
                this.$refs.pagination.currentPage = 1;
            }

        },

        resetPage() {
            this.resetting = true;

            if (this.$refs.pagination.currentPage !== 1) { // Access childs current page, changing this to 1 will trigger a data reload.
                this.$refs.pagination.currentPage = 1;
            }
            this.retrievePlanets(1, 'All', 'All');
            this.$router.push(this.$route.path).catch(() => {}); // Catch the error because we dont care about redirecting to same page.
        },

        resetPageValue() {
            return true;
        },

        handlePageChange(page) {
            if (this.resetting || this.searching) {
                return;
            }

            const name = typeof this.$route.query.name == "undefined" ? "All" : this.$route.query.name;
            this.$router.push({
                query: Object.assign({}, this.$route.query, {
                    page: page,
                    type: this.type,
                    name: name
                })
            });
            this.retrievePlanets(page, this.type, name);
        }
    },
    watch: { // Watch for data change in which page the user is currently on, call API to get new data when it changes.
        "type": function () {
            if (this.resetting || this.searching) {
                return;
            }

            if (this.$refs.pagination.currentPage !== 1) { // Access childs current page, changing this to 1 will trigger a data reload.
                this.$refs.pagination.currentPage = 1;
            } else {
                this.$router.push({
                    query: Object.assign({}, this.$route.query, {
                        page: 1,
                        type: this.type,
                        name: "All"
                    })
                });
                this.retrievePlanets(1, this.type, "All");
            }
        },

        "planets": function (val) {
            if (val === null) {
                return;
            }

            if (val.length === 1) {
                this.currentPlanet = val[0];
                this.currentIndex = 0;
            }
        }
    },
    mounted() {
        this.retrievePlanetTypes();
        const name = typeof this.$route.query.name == "undefined" || this.$route.query.name == "All" ? "" : this.$route.query.name
        const page = typeof this.$route.query.page == "undefined" ? 1 : parseInt(this.$route.query.page)
        this.retrievePlanets(page, this.type, name);
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
    transition: all 1s ease;
}

.slide-fade-leave-active {
    transition: all 1s ease;
}

.slide-fade-enter,
.slide-fade-leave-to

/* .slide-fade-leave-active below version 2.1.8 */
    {
    transform: translatey(300px);
    opacity: 0;
}
</style>
