Rails.application.routes.draw do
  resources :quotes
  resources :users
  resources :buildings
  resource :dashboard
  resource :session

  root "welcome#index"
end
