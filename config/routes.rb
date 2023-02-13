Rails.application.routes.draw do
  resources :quotes

  root "welcome#index"
end
