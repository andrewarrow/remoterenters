Rails.application.routes.draw do
  resources :quotes
  resources :users
  resource :dashboard
  resource :session

  root "welcome#index"
end
